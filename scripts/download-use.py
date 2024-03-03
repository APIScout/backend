import tensorflow_hub as hub
import tensorflow.compat.v1 as tf


MODEL_NAME = 'universal-encoder'
VERSION = 1
SERVE_PATH = './models/{}/{}'.format(MODEL_NAME, VERSION)


with tf.Graph().as_default():
    module = hub.Module("https://tfhub.dev/google/universal-sentence-encoder/2")
    text = tf.placeholder(tf.string, [None])
    embedding = module(text)

    init_op = tf.group([tf.global_variables_initializer(),
                        tf.tables_initializer()])

    with tf.Session() as session:
        session.run(init_op)
        tf.saved_model.simple_save(
            session,
            SERVE_PATH,
            inputs={"text": text},
            outputs={"embedding": embedding},
            legacy_init_op=tf.tables_initializer()
        )
